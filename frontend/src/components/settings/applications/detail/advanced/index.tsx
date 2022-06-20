import style from './index.less';
import { Button, Popconfirm } from 'antd';
import { useDispatch } from 'umi';
const Advanced: React.FC = () => {
  const dispatch = useDispatch();
  const onDelete = () => {
    dispatch({
      type: 'application/delete',
    });
  };
  return (
    <div className={style.advanced}>
      <span className={style.title}>危险区</span>
      <div className={style.dangerZone}>
        <div className={style.dangerLine}>
          <span className={style.dangerTitle}>删除此应用程序</span>
          <span className={style.danger}>这不能被撤销。请确定。</span>
        </div>
        <Popconfirm
          placement="topRight"
          title="删除应用程序将使用户授权的任何访问令牌无效。您确定要删除此应用程序吗？"
          onConfirm={onDelete}
          okText="确定"
          cancelText="取消"
        >
          <Button className={style.delete} size="large" danger>
            删除应用程序
          </Button>
        </Popconfirm>
      </div>
    </div>
  );
};
export default Advanced;
