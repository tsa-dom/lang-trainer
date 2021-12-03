import React, { useEffect } from 'react'
import { useDispatch, useSelector } from 'react-redux'
import { Route, Switch } from 'react-router-dom'
import './App.css'
import AppBar from './components/AppBar'
import GroupPage from './components/GroupPage'
import Groups from './components/Groups'
import LoginForm from './components/LoginForm'
import MainPage from './components/MainPage'
import Notification from './components/Styled/Notification'
import { setUser } from './features/userSlice'
import useUser from './hooks/user'

const App = () => {
  const dispatch = useDispatch()
  const user = useSelector(state => state.users.currentUser)
  const { message, type } = useSelector(state => state.notifications)
  const { authorize } = useUser()

  useEffect(async () => {
    const token = localStorage.getItem('app-token')
    if (token) {
      const user = await authorize(token)
      if (user && !user.errors) dispatch(setUser(user))
      else localStorage.removeItem('app-token')
    }
  }, [])

  return (
    <>
      <AppBar />
      <Notification message={message} type={type} />
      <div className="page-container">
        <Switch>
          <Route path="/practice">
            {user &&
              <div>Hello world</div>
            }
          </Route>
          <Route path="/groups">
            {user &&
              <Groups />
            }
          </Route>
          <Route path="/group">
            {user &&
              <GroupPage />
            }
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
