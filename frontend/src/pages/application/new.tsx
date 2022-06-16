import styles from './new.less';
import { Layout } from 'antd';
import { NewApplication } from '@/components';
const { Content } = Layout;

const New: React.FC = () => {
  return (
    <Layout className={styles.application}>
      <Content className={styles.main}>
        <NewApplication />
      </Content>
    </Layout>
  );
};
export default New;
