import style from './index.less';
import type { IRouteComponentProps } from 'umi';
import { Layout } from 'antd';
import { history } from 'umi';
const ApplicationLayout: React.FC<IRouteComponentProps> = ({ children }) => {
  if (history.location.pathname.startsWith('/login')) {
    return (
    <Layout className={style.layout}>
    <Layout.Content className={style.content}>{children}</Layout.Content>
  </Layout>
    )
  }
  return (
    <Layout className={style.layout}>
      <Layout.Header />
      <Layout.Content className={style.content}>{children}</Layout.Content>
    </Layout>
  );
};
export default ApplicationLayout;
