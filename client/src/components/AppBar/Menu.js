import React from 'react'
import { useTranslation } from 'react-i18next'
import { useDispatch, useSelector } from 'react-redux'
import { useHistory } from 'react-router'
import { clearUser, setLanguage } from '../../features/userSlice'
import Button from '../Styled/Button'
import { resources } from '../../i18n'
import { FaUserCog, FaUserGraduate, FaUserEdit } from 'react-icons/fa'

const Menu = () => {
  const { t } = useTranslation()
  const user = useSelector(state => state.users.currentUser)
  const language = useSelector(state => state.users.language)
  const dispatch = useDispatch()
  const history = useHistory()

  const loginPage = () => {
    history.push('/login')
  }

  const groupPage = () => {
    history.push('/groups')
  }

  const practisePage = () => {
    history.push('/practice')
  }

  const handleLogout = () => {
    localStorage.removeItem('app-token')
    dispatch(clearUser())
    history.push('/login')
  }

  const handleChangeLanguage = (value) => {
    const language = Object.keys(resources.fi.translation).find(key => resources.fi.translation[key] === value)
    dispatch(setLanguage(language))
  }

  return (
    <div className="appbar-menu">
      {!user &&
        <Button className="appbar-button" text={t('menu-login')} onClick={loginPage} />
      }
      {user &&
        <>
          <Button className="appbar-button" text={t('menu-practise')} onClick={practisePage} />
          <Button className="appbar-button" text={t('menu-groups')} onClick={groupPage} />
          <Button className="appbar-button" text={t('menu-logout')} onClick={handleLogout} />
        </>
      }
      <Button className="appbar-button"
        style={{ position: 'relative' }}
        dropdown
        options={Object.keys(resources).map(key => t(resources[key].language))}
        text={t(language)}
        onClick={handleChangeLanguage}
      />
      {user &&
        <Button
          className="appbar-button"
          text={<span>
            <span style={{ marginRight: 5 }}>{user.username}</span>
            {user.privileges === 'admin' && <FaUserEdit style={{ width: 23, marginBottom: -2 }} />}
            {user.privileges === 'teacher' && <FaUserGraduate style={{ width: 23, marginBottom: -2 }} />}
            {user.privileges === 'student' && <FaUserCog style={{ width: 23, marginBottom: -2 }} />}
          </span>}
        />
      }
    </div>
  )
}

export default Menu