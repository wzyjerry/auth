import style from './index.less';
import { Tabs } from 'antd';
import DetailTab from './detail';
interface DetailProp {
  id: string
}
const Detail: React.FC<DetailProp> = (prop) => {
  return (
    <Tabs className={style.tab} tabPosition='left'>
      <Tabs.TabPane tab='详情' key='1'>
        <DetailTab id={prop.id}></DetailTab>
      </Tabs.TabPane>
      <Tabs.TabPane tab='高级' key='2'>

      </Tabs.TabPane>
    </Tabs>
  )
}

export default Detail;
