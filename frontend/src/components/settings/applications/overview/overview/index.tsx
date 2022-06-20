import style from './index.less';
import type { ApplicationOverview } from '@/api/application/v1/application';
import { Avatar, Button } from 'antd';
import { history } from 'umi';
interface OverviewProp {
  overview: ApplicationOverview;
}
const Overview: React.FC<OverviewProp> = ({ overview }) => {
  const detail = () => {
    history.push(`/settings/applications/${overview.id}`);
  };
  return (
    <div className={style.overview}>
      <Button className={style.logo} type="link" onClick={detail}>
        <Avatar src={overview.logo} size={56}>
          {overview.name[0]}
        </Avatar>
      </Button>
      <div className={style.infoArea}>
        <Button className={style.name} type="link" onClick={detail}>
          {overview.name}
        </Button>
        <span className={style.description}>{overview.maskedDescription}</span>
      </div>
    </div>
  );
};
export default Overview;
