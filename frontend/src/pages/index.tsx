import { LoginNavbar } from '@/components';
import styles from './index.less';

const App = () => {
  return (
    <div>
      <h1 className={styles.title}>Auth frontend</h1>
      <LoginNavbar></LoginNavbar>
    </div>
  );
};

export default App;
