import { ColumnsType } from 'antd/es/table'
import { PortUsage } from '../../types/PortUsage'

export const tableColumns: ColumnsType<PortUsage> = [
  {
    title: '协议',
    dataIndex: 'protocol',
    key: 'protocol',
    sorter: (a, b) => a.protocol.localeCompare(b.protocol),
  },
  {
    title: '本地地址',
    dataIndex: 'localAddr',
    key: 'localAddr',
  },
  {
    title: '端口',
    dataIndex: 'port',
    key: 'port',
    sorter: (a, b) => Number(a.port) - Number(b.port),
  },
  {
    title: '状态',
    dataIndex: 'state',
    key: 'state',
  },
  {
    title: '进程ID',
    dataIndex: 'pid',
    key: 'pid',
  },
  {
    title: '程序名称',
    dataIndex: 'program',
    key: 'program',
  },
]