import { useEffect, useState } from 'react'
import { PortUsage } from '../types/PortUsage'

export const usePortData = () => {
	const [data, setData] = useState<PortUsage[]>([])
	const [loading, setLoading] = useState(false)
	const [error, setError] = useState<string | null>(null)

	const fetchData = async () => {
		try {
			setLoading(true)
			const response = await fetch('/api/list')
			const result = await response.json()
			setData(result.data)
			setError(null)
		} catch (err) {
			setError('获取数据失败')
			console.error('Failed to fetch port data:', err)
		} finally {
			setLoading(false)
		}
	}

	useEffect(() => {
		fetchData()
		const interval = setInterval(fetchData, 5000)
		return () => clearInterval(interval)
	}, [])

	return { data, loading, error, refetch: fetchData }
}