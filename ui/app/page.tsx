'use client';

import { useEffect } from "react";
import Head from "next/head";

export default function Home() {
  useEffect(() => {
    async function fetchGraph() {
      try {
        const response = await fetch("http://localhost:8080/graph"); // Use an API route for fetching data
        const data = await response.json();

        // Map nodes
        const nodes = data.nodes.map((node) => ({
          id: node.id,
          label: `${node.id}\n(${node.type})`,
        }));

        // Map edges
        const edges = [];
        data.nodes.forEach((node) => {
          (node.depends || []).forEach((dep) => {
            edges.push({ from: dep, to: node.id });
          });
        });

        // Initialize the network graph
        const container = document.getElementById("network");
        const networkData = {
          nodes: new vis.DataSet(nodes),
          edges: new vis.DataSet(edges),
        };

        const options = {
          nodes: {
            shape: "box",
            font: { align: "center" },
          },
          edges: {
            arrows: "to",
          },
        };

        new vis.Network(container, networkData, options);
      } catch (error) {
        console.error("Error fetching or visualizing graph data:", error);
        alert("Failed to load the graph. Check the console for details.");
      }
    }

    fetchGraph();
  }, []);

  return (
    <>
      <Head>
        <title>KubeMap - Terraform State Visualizer</title>
        <script src="https://cdnjs.cloudflare.com/ajax/libs/vis/4.21.0/vis.min.js"></script>
        <link
          href="https://cdnjs.cloudflare.com/ajax/libs/vis/4.21.0/vis.min.css"
          rel="stylesheet"
        />
      </Head>
      <div style={{ display: "flex", height: "100vh", margin: 0, padding: 0 }}>
        <div id="sidebar" style={sidebarStyle}>
          <h2>KubeMap</h2>
          <ul style={ulStyle}>
            <li onClick={() => alert("Load New Data (Placeholder)")} style={liStyle}>
              Load Data
            </li>
            <li onClick={() => alert("Refresh Graph (Placeholder)")} style={liStyle}>
              Refresh Graph
            </li>
            <li onClick={() => alert("Export Graph (Placeholder)")} style={liStyle}>
              Export
            </li>
            <li onClick={() => alert("Help (Placeholder)")} style={liStyle}>
              Help
            </li>
          </ul>
        </div>
        <div id="content" style={contentStyle}>
          <header style={headerStyle}>
            <h1 style={{ margin: 0, fontSize: "24px" }}>
              Terraform Resource Dependency Graph
            </h1>
          </header>
          <div id="network-container" style={networkContainerStyle}>
            <div id="network" style={networkStyle}></div>
          </div>
        </div>
      </div>
      <footer style={footerStyle}>
        Built with ❤️ by{" "}
        <a href="https://www.linkedin.com/in/dhyanio/" target="_blank" style={footerLinkStyle}>
          Dhyanio
        </a>
      </footer>
    </>
  );
}

const sidebarStyle = {
  width: "250px",
  backgroundColor: "#333",
  color: "white",
  display: "flex",
  flexDirection: "column",
  boxShadow: "2px 0 5px rgba(0, 0, 0, 0.1)",
};

const ulStyle = {
  listStyle: "none",
  padding: 0,
  margin: 0,
};

const liStyle = {
  padding: "10px 15px",
  borderBottom: "1px solid #555",
  cursor: "pointer",
  textAlign: "left",
};

const headerStyle = {
  backgroundColor: "#4CAF50",
  color: "white",
  padding: "15px",
  textAlign: "center",
  boxShadow: "0px 2px 5px rgba(0, 0, 0, 0.1)",
  position: "fixed",
  top: 0,
  left: "250px",
  right: 0,
  zIndex: 1000,
};

const contentStyle = {
  marginLeft: "250px",
  marginTop: "80px",
  flex: 1,
  display: "flex",
  flexDirection: "column",
  alignItems: "center",
  justifyContent: "center",
};

const networkContainerStyle = {
  width: "100%",
  height: "calc(100vh - 120px)",
  padding: "20px",
};

const networkStyle = {
  width: "100%",
  height: "100%",
  border: "1px solid #ddd",
  backgroundColor: "#fff",
  boxShadow: "0px 2px 10px rgba(0, 0, 0, 0.1)",
};

const footerStyle = {
  textAlign: "center",
  padding: "10px 0",
  backgroundColor: "#f1f1f1",
  borderTop: "1px solid #ddd",
};

const footerLinkStyle = {
  color: "#4CAF50",
  textDecoration: "none",
  fontWeight: "bold",
};
