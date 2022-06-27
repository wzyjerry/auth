import styles from './index.less';
import { useSelector } from 'umi';
import type { OAuth, User } from 'umi';
import { Avatar, Button } from 'antd';
import { GlobalOutlined } from '@ant-design/icons';
const Authorize: React.FC = () => {
  const { oauth, user } = useSelector(
    ({ login_oauth2_authorize, user }: { login_oauth2_authorize: OAuth; user: User }) => ({
      oauth: login_oauth2_authorize,
      user,
    }),
  );
  console.log(user);
  return (
    <div className={styles.authorize}>
      <Avatar size="large" src={oauth.logo} shape="square">
        oauth.name[0]
      </Avatar>
      <span className={styles.title}>授权 {oauth.name}</span>
      <div className={styles.info}>
        <div className={styles.infoBody}>
          <div className={styles.infoLine}>
            <div className={styles.iconArea}>
              <Avatar size="large" src={user.avatar}>
                user.nickname[0]
              </Avatar>
            </div>
            <div className={styles.infoArea}>
              <div className={styles.infoLabel}>
                <Button type="link" className={styles.owner}>
                  wzyjerry
                </Button>
                <span>的</span>
                <span className={styles.name}> {oauth.name} </span>
                <span>服务</span>
              </div>
              <div className={styles.infoDescription}>
                <span>想要访问您的</span>
                <span className={styles.nickname}> {user.nickname} </span>
                <span>账户</span>
              </div>
            </div>
          </div>
          <div className={styles.infoLine}>
            <div className={styles.iconArea}>
              <GlobalOutlined />
            </div>
            <div className={styles.infoArea}>
              <div className={styles.infoLabel}>
                <span className={styles.data}>仅公开数据</span>
              </div>
              <div className={styles.infoDescription}>
                <span>此应用程序将能够识别您并获取公共信息</span>
              </div>
            </div>
          </div>
        </div>
        <div className={styles.infoFooter}>
          <div className={styles.actions}>
            <Button className={styles.cancel} size="large">
              取消
            </Button>
            <Button className={styles.submit} size="large" type="primary">
              授权 {user.nickname}
            </Button>
          </div>
          <div className={styles.footerInfo}>
            <span className={styles.footerLabel}>授权将重定向到</span>
            <span className={styles.homepage}>{oauth.homepage}</span>
          </div>
        </div>
      </div>
    </div>
  );
};
export default Authorize;
