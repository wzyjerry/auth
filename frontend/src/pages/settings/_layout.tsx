import style from './_layout.less';
import type { IRouteComponentProps } from 'umi';
import { SettingNavbar } from '@/components';
const ApplicationLayout: React.FC<IRouteComponentProps> = ({ children }) => {
  return (
    <div className={style.settings}>
      <SettingNavbar />
      {children}
    </div>
  );
};
export default ApplicationLayout;
