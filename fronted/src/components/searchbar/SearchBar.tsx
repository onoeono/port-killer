import { Input, Space, Select } from 'antd'
import { SearchOutlined } from '@ant-design/icons'

const { Option } = Select

interface PortSearchBarProps {
  onSearch: (value: string) => void
  onProtocolFilter: (value: string) => void
}

const PortSearchBar = ({ onSearch, onProtocolFilter }: PortSearchBarProps) => {
  return (
    <Space style={{ marginBottom: 16 }}>
      <Input
        placeholder="搜索端口或程序名称"
        prefix={<SearchOutlined />}
        onChange={(e) => onSearch(e.target.value)}
        style={{ width: 200 }}
      />
      <Select
        defaultValue="all"
        style={{ width: 120 }}
        onChange={onProtocolFilter}
      >
        <Option value="all">全部协议</Option>
        <Option value="tcp">TCP</Option>
        <Option value="udp">UDP</Option>
      </Select>
    </Space>
  )
}

export default PortSearchBar