import style from './index.less';
import { useSelector } from 'umi';
import type { GetAllReply } from '@/api/application/v1/application';
import { Button } from 'antd';
import OverviewItem from './overview';
import { history } from 'umi';
const Overview: React.FC = () => {
  const overviews = useSelector(
    ({ application }: { application: GetAllReply }) => application.applicationOverviews,
  );
  const newApplication = () => {
    history.push('/settings/applications/new');
  };
  const renderOverviews = () => {
    if (overviews.length > 0) {
      return (
        <div className={style.overviewList}>
          {overviews.map((overview) => (
            <OverviewItem key={overview.id} overview={overview} />
          ))}
        </div>
      );
    } else {
      return <span>创建您的第一个应用程序。</span>;
    }
  };
  return (
    <div className={style.overview}>
      <div className={style.titleLine}>
        <span className={style.title}>应用程序</span>
        <Button className={style.new} type="primary" size="large" onClick={newApplication}>
          创建应用程序
        </Button>
      </div>
      <hr className={style.line} />
      {renderOverviews()}
    </div>
  );
};
export default Overview;
