import React from 'react'
import { useTranslation } from 'react-i18next'
import { useHistory } from 'react-router'
import Button from '../Styled/Button'

const Menu = ({ currentUser, setCurrentUser }) => {
  const { t } = useTranslation('translation')
  const history = useHistory()

  const loginPage = () => {
    history.push('/login')
  }

  const handleLogout = () => {
    setCurrentUser(undefined)
    localStorage.removeItem('app-token')
  }

  return (
    <div className="appbar-menu">
      {!currentUser &&
        <Button className="appbar-button" text={t('menu-login')} onClick={loginPage} />
      }
      {currentUser &&
        <Button className="appbar-button" text={t('menu-logout')} onClick={handleLogout} />
      }
    </div>
  )
}

export default Menu