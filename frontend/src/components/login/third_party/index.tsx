import style from './index.less';
import Github from './github';
import Microsoft from './microsoft';

interface ThirdPartyProp {
  className?: string;
  returnTo: string;
}
const ThirdParty: React.FC<ThirdPartyProp> = ({ className, returnTo }) => {
  return (
    <div className={className}>
      <span className={style.title}>其他方式登录</span>
      <div className={style.list}>
        <Github className={style.item} returnTo={returnTo} />
        <Microsoft className={style.item} returnTo={returnTo} />
      </div>
    </div>
  );
};
export default ThirdParty;
