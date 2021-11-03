import React from 'react'
import { useTranslation } from 'react-i18next'
import { useDispatch, useSelector } from 'react-redux'
import { useHistory } from 'react-router'
import { clearUser } from '../../features/userSlice'
import Button from '../Styled/Button'

const Menu = () => {
  const { t } = useTranslation('translation')
  const user = useSelector(state => state.users.currentUser)
  const dispatch = useDispatch()
  const history = useHistory()

  const loginPage = () => {
    history.push('/login')
  }

  const groupPage = () => {
    history.push('/groups')
  }

  const handleLogout = () => {
    localStorage.removeItem('app-token')
    dispatch(clearUser())
    history.push('/login')
  }

  return (
    <div className="appbar-menu">
      {!user &&
        <Button className="appbar-button" text={t('menu-login')} onClick={loginPage} />
      }
      {user &&
        <>
          <Button className="appbar-button" text={t('menu-groups')} onClick={groupPage} />
          <Button className="appbar-button" text={t('menu-logout')} onClick={handleLogout} />
        </>
      }
    </div>
  )
}

export default Menu