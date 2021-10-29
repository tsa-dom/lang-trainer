import React, { useEffect, useState } from 'react'
import './App.css'
import AppBar from './components/AppBar'
import PageContainer from './components/Page'
import usePing from './hooks/ping'

const App = () => {
  const { getPing, result, loading } = usePing()
  const [page, setPage] = useState('first')
  const [, setMessage] = useState()

  useEffect(() => getPing(), [])

  useEffect(() => {
    if (!loading && !result.errors) setMessage(result.message)
  }, [loading])

  return (
    <>
      <AppBar setPage={setPage} />
      <PageContainer page={page} />
    </>
  )
}

export default App
