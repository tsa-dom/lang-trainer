import React, { useEffect } from 'react'
import { useDispatch } from 'react-redux'
import { Route, Switch } from 'react-router-dom'
import './App.css'
import AppBar from './components/AppBar'
import GroupPage from './components/GroupPage'
import Groups from './components/Groups'
import LoginForm from './components/LoginForm'
import MainPage from './components/MainPage'
import { setUser } from './features/userSlice'
import useUser from './hooks/user'

const App = () => {
  const dispatch = useDispatch()
  const { authorize, user } = useUser()

  useEffect(() => {
    const token = localStorage.getItem('app-token')
    if (token) {
      authorize(token)
    }
  }, [])

  useEffect(() => {
    if (user) dispatch(setUser(user))
  }, [user])

  return (
    <>
      <AppBar />
      <div className="page-container">
        <Switch>
          <Route path="/groups">
            <Groups />
          </Route>
          <Route path="/group">
            <GroupPage />
          </Route>
          <Route path="/login">
            <LoginForm />
          </Route>
          <Route path="/">
            <MainPage />
          </Route>
        </Switch>
      </div>
    </>
  )
}

export default App
