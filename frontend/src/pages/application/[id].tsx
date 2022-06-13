import styles from './new.less';
import { Layout } from 'antd';
import { ApplicationDetail } from '@/components';
import { useParams } from 'umi';
const { Content } = Layout;

interface Params {
  id: string
}

const Detail: React.FC = () => {
  const { id } = useParams<Params>();
  return (
    <Layout className={styles.application}>
      <Content className={styles.main}>
        <ApplicationDetail id={id}></ApplicationDetail>
      </Content>
    </Layout>
  );
};
export default Detail;
