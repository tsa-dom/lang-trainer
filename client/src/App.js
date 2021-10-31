import React, { useEffect, useState } from 'react'
import { Route, Switch } from 'react-router-dom'
import './App.css'
import AppBar from './components/AppBar'
import LoginForm from './components/LoginForm'
import useUser from './hooks/user'

const App = () => {
  const [currentUser, setCurrentUser] = useState()
  const { authorize, user } = useUser()

  useEffect(() => {
    const token = localStorage.getItem('app-token')
    if (token) {
      authorize(token)
    }
  }, [])

  useEffect(() => {
    if (user) setCurrentUser(user)
  }, [user])

  return (
    <>
      <AppBar
        currentUser={currentUser}
        setCurrentUser={setCurrentUser}
      />
      <div className="page-container">
        <Switch>
          <Route path="/login">
            <LoginForm setCurrentUser={setCurrentUser}/>
          </Route>
          <Route path="/">
            <div>hello world</div>
          </Route>
        </Switch>
      </div>
    </>
  )
}

export default App
