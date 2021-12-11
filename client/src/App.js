import React, { useEffect } from 'react'
import { useSelector } from 'react-redux'
import { Route, Switch } from 'react-router-dom'
import './App.css'
import AppBar from './components/AppBar'
import GroupPage from './components/GroupPage'
import Groups from './components/Groups'
import LoginForm from './components/LoginForm'
import MainPage from './components/MainPage'
import Notification from './components/Styled/Notification'
import { useTranslation } from 'react-i18next'
import { fetchUser } from './utils/fetcher'

const App = () => {
  const user = useSelector(state => state.users.currentUser)
  const language = useSelector(state => state.users.language)
  const { message, type } = useSelector(state => state.notifications)
  const { i18n } = useTranslation()

  useEffect(() => i18n.changeLanguage(language), [language])

  useEffect(fetchUser, [])

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
