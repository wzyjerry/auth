import styles from './new.less';
import { Layout } from 'antd';
import { ApplicationDetail } from '@/components';
import { useDispatch, useParams } from 'umi';
import { useEffect } from 'react';

const { Content } = Layout;

interface Params {
  id: string
}

const Detail: React.FC = () => {
  const dispatch = useDispatch();
  const { id } = useParams<Params>();
  useEffect(() => {
    dispatch({
      type: 'application/setup',
      payload: { id },
    })
  }, [id])
  return (
    <Layout className={styles.application}>
      <Content className={styles.main}>
        <ApplicationDetail></ApplicationDetail>
      </Content>
    </Layout>
  );
};
export default Detail;
