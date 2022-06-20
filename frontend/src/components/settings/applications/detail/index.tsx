import style from './index.less';
import { Tabs } from 'antd';
import DetailTab from './detail';
import AdvancedTab from './advanced';
const Detail: React.FC = () => {
  return (
    <Tabs className={style.tab} tabPosition="left">
      <Tabs.TabPane tab="详情" key="1">
        <DetailTab />
      </Tabs.TabPane>
      <Tabs.TabPane tab="高级" key="2">
        <AdvancedTab />
      </Tabs.TabPane>
    </Tabs>
  );
};

export default Detail;
