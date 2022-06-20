import style from './index.less';
import type { IRouteComponentProps } from 'umi';
import { Layout } from 'antd';
const ApplicationLayout: React.FC<IRouteComponentProps> = ({ children }) => {
  return (
    <Layout className={style.layout}>
      <Layout.Header />
      <Layout.Content className={style.content}>{children}</Layout.Content>
    </Layout>
  );
};
export default ApplicationLayout;
