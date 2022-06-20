import style from './index.less';
import { Breadcrumb } from 'antd';
import { withRouter, Link, useSelector } from 'umi';
import type { Application } from 'umi';
const SettingNavbar = withRouter((prop) => {
  const { location } = prop;
  const pathSnippets = location.pathname.split('/').filter((i) => i);
  const state = useSelector((state: { application: Application }) => state);
  const breadcrumbNameMap = (slice: string[]) => {
    if (slice.length >= 1 && slice[0] == 'settings') {
      if (slice.length >= 2 && slice[1] == 'applications') {
        if (slice.length == 3) {
          switch (slice[2]) {
            case 'new':
              return '创建应用程序';
            default:
              const { application } = state;
              return application.name;
          }
        }
        return '应用程序';
      }
      return '设置';
    }
  };
  const extraBreadcrumbItems = pathSnippets.map((_, index) => {
    const url = `/${pathSnippets.slice(0, index + 1).join('/')}`;
    if (index + 1 == pathSnippets.length) {
      return (
        <Breadcrumb.Item key={url}>
          {breadcrumbNameMap(pathSnippets.slice(0, index + 1))}
        </Breadcrumb.Item>
      );
    }
    return (
      <Breadcrumb.Item key={url}>
        <Link to={url}>{breadcrumbNameMap(pathSnippets.slice(0, index + 1))}</Link>
      </Breadcrumb.Item>
    );
  });
  const breadcrumbItems = [
    <Breadcrumb.Item key="/">
      <Link to="/">主页</Link>
    </Breadcrumb.Item>,
  ].concat(extraBreadcrumbItems);
  return <Breadcrumb className={style.navbar}>{breadcrumbItems}</Breadcrumb>;
});
export default SettingNavbar;
