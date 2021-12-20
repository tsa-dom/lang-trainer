import React, { useEffect } from 'react'
import { useSelector } from 'react-redux'
import { Route, Routes } from 'react-router-dom'
import './App.css'
import AppBar from './components/AppBar'
import { useTranslation } from 'react-i18next'
import Login from './components/Login'
import Groups from './components/Groups'
import Group from './components/Group'
import useFetch from './hooks/fetcher'
import Templates from './components/Templates/List'

const App = () => {
  const user = useSelector(state => state.users.currentUser)
  const language = useSelector(state => state.users.language)
  const { i18n } = useTranslation()
  const { fetchUser } = useFetch()

  useEffect(() => i18n.changeLanguage(language), [language, i18n])

  useEffect(fetchUser, [])

  return (
    <>
      <AppBar />
      <div>
        <Routes>
          {user &&
          <>
            <Route path="/groups" element={<Groups />} />
            <Route path="/templates" element={<Templates />} />
            <Route path="/group" element={<Group />} />
          </>
          }
          <Route path="/login" element={<Login />} />
          <Route path="/" element={<div>Main page</div>} />
          <Route path="*" element={<>404 Not Found</>} />
        </Routes>
      </div>
    </>
  )
}

export default App
