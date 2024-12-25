import { useEffect } from 'react';
import { Network } from 'vis-network/standalone/umd/vis-network.min';

export default function NetworkGraph() {
  useEffect(() => {
    async function fetchGraph() {
      try {
        const response = await fetch('/api/graph');
        const data = await response.json();

        const nodes = data.nodes.map((node) => ({
          id: node.id,
          label: `${node.id}\n(${node.type})`,
        }));

        const edges = [];
        data.nodes.forEach((node) => {
          (node.depends || []).forEach((dep) => {
            edges.push({ from: dep, to: node.id });
          });
        });

        const container = document.getElementById('network');
        const networkData = {
          nodes: new Network.DataSet(nodes),
          edges: new Network.DataSet(edges),
        };

        const options = {
          nodes: {
            shape: 'box',
            font: { align: 'center' },
          },
          edges: {
            arrows: 'to',
          },
        };

        new Network(container, networkData, options);
      } catch (error) {
        console.error('Error fetching or visualizing graph data:', error);
      }
    }

    fetchGraph();
  }, []);

  return <div id="network" style={{ height: '100%', width: '100%' }} />;
}
