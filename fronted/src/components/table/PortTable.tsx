import { Table } from 'antd'
import { tableColumns } from './TableColumn'
import { PortUsage } from '../../types/PortUsage'

interface PortTableProps {
  loading: boolean
  data: PortUsage[]
}

const PortTable = ({ loading, data }: PortTableProps) => {
  return (
    <Table
      columns={tableColumns}
      dataSource={data.data}
      loading={loading}
      rowKey={(record) => `${record.protocol}-${record.port}-${record.pid}`}
      pagination={{
        showSizeChanger: true,
        showQuickJumper: true,
        showTotal: (total) => `共 ${total} 条`,
      }}
    />
  )
}

export default PortTable