import style from './index.less';
import Github from './github';
import Microsoft from './microsoft';

interface ThirdPartyProp {
  className?: string
  returnTo: string
}
const ThirdParty: React.FC<ThirdPartyProp> = (prop) => {
  console.log(prop.returnTo)
  return (
    <div className={prop.className}>
      <span className={style.title}>其他方式登录</span>
      <div className={style.list}>
        <Github className={style.item} returnTo={prop.returnTo}></Github>
        <Microsoft className={style.item} returnTo={prop.returnTo}></Microsoft>
      </div>
    </div>
  )
};
export default ThirdParty;
