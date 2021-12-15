import React, { useEffect } from 'react'
import { useSelector } from 'react-redux'
import { Route, Routes } from 'react-router-dom'
import './App.css'
import AppBar from './components/AppBar'
//import GroupPage from './components/GroupPage'
//import Groups from './components/Groups'
//import LoginForm from './components/LoginForm'
//import MainPage from './components/MainPage'
//import Notification from './components/Styled/Notification'
import { useTranslation } from 'react-i18next'
import Login from './components/Login'
//import { fetchUser } from './utils/fetcher'
//import Login from './components/Login'

const App = () => {
  //const user = useSelector(state => state.users.currentUser)
  const language = useSelector(state => state.users.language)
  const { i18n } = useTranslation()

  useEffect(() => i18n.changeLanguage(language), [language, i18n])

  //useEffect(fetchUser, [])

  return (
    <>
      <AppBar />
      <div>
        <Routes>
          <Route path="/login" element={<Login />} />
          <Route path="/"  element={<></>} />
        </Routes>
      </div>
    </>
  )
}

export default App
