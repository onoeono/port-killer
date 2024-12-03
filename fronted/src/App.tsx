import { useState } from 'react'
import { Layout, Typography, Alert } from 'antd'
import PortTable from './components/table/PortTable'
import PortSearchBar from './components/searchbar/SearchBar'
import { usePortData } from './hooks/usePortData'
import { PortUsage } from './types/PortUsage'

const { Header, Content } = Layout
const { Title } = Typography

function App() {
  const { data, loading, error } = usePortData()
  const [searchText, setSearchText] = useState('')
  const [protocol, setProtocol] = useState('all')

  // const filteredData = data.filter((item: PortUsage) => {
  //   const matchesSearch =
  //     item.port.includes(searchText) ||
  //     item.program.toLowerCase().includes(searchText.toLowerCase())
  //   const matchesProtocol =
  //     protocol === 'all' || item.protocol.toLowerCase() === protocol
  //
  //   return matchesSearch && matchesProtocol
  // })

  return (
    <Layout style={{ minHeight: '100vh' }}>
      <Header style={{ background: '#fff', padding: '0 16px' }}>
        <Title level={3}>端口使用情况监控</Title>
      </Header>
      <Content style={{ padding: '24px' }}>
        {error && (
          <Alert
            message="错误"
            description={error}
            type="error"
            showIcon
            style={{ marginBottom: 16 }}
          />
        )}
        <PortSearchBar
          onSearch={setSearchText}
          onProtocolFilter={setProtocol}
        />
        <PortTable loading={loading} data={data} />
      </Content>
    </Layout>
  )
}

export default App