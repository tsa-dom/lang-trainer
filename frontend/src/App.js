/* eslint-disable no-unused-vars */
import React, { useEffect, useState } from 'react'
import './App.css'
import usePing from './hooks/ping'

const App = () => {
  const { getPing, result, loading } = usePing()
  const [message, setMessage] = useState()

  useEffect(() => getPing(), [])

  useEffect(() => {
    if (!loading && !result.errors) setMessage(result.message)
  }, [loading])

  return (
    <div>{message}</div>
  )
}

export default App
