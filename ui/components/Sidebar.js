import styles from '../styles/Sidebar.module.css';

export default function Sidebar() {
  return (
    <div className={styles.sidebar}>
      <h2>KubeMap</h2>
      <ul>
        <li onClick={() => alert('Load New Data (Placeholder)')}>Load Data</li>
        <li onClick={() => alert('Refresh Graph (Placeholder)')}>Refresh Graph</li>
        <li onClick={() => alert('Export Graph (Placeholder)')}>Export</li>
        <li onClick={() => alert('Help (Placeholder)')}>Help</li>
      </ul>
    </div>
  );
}
