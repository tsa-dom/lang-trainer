import React from 'react'
import { Container, Button, Navbar, NavDropdown } from 'react-bootstrap'
import { useTranslation } from 'react-i18next'
import { useSelector } from 'react-redux'
import { useNavigate } from 'react-router-dom'
import { resources } from '../../i18n'
import { clearUser, setLanguage } from '../../features/userSlice'
import { useDispatch } from 'react-redux'

const AppBar = () => {
  const { t } = useTranslation()
  const navigate = useNavigate()
  const language = useSelector(state => state.users.language)
  const dispatch = useDispatch()
  const user = useSelector(state => state.users.currentUser)

  const languages = Object.keys(resources).map(key => t(resources[key].language))

  const handleChangeLanguage = (value) => {
    const language = Object.keys(resources.fi.translation).find(key => resources.fi.translation[key] === value)
    dispatch(setLanguage(language))
  }

  const handleLogout = () => {
    localStorage.removeItem('app-token')
    dispatch(clearUser())
    navigate('/login')
  }

  return (
    <Navbar className="nav-menu" bg="dark" variant="dark" expand="lg">
      <Container>
        <Navbar.Brand><span className="nav-menu-brand">Lang Trainer</span></Navbar.Brand>
        <Navbar.Toggle />
        <Navbar.Collapse className="justify-content-end">
          <NavDropdown title={t(language)}>
            {languages.map(r => <NavDropdown.Item onClick={() => handleChangeLanguage(r)} key={r}>{t(r)}</NavDropdown.Item>)}
          </NavDropdown>
          {user &&
            <NavDropdown style={{ marginRight: 20 }} title={t('groups')}>
              <NavDropdown.Item onClick={() => navigate('/groups')}>{t('group-list')}</NavDropdown.Item>
              <NavDropdown.Item onClick={() => navigate('/templates')}>{t('templates')}</NavDropdown.Item>
            </NavDropdown>
          }
          {!user &&
            <Button onClick={() => navigate('/login')} className="button-menu">{t('login')}</Button>
          }
          {user &&
            <Button onClick={handleLogout} className="button-menu">{t('logout')}</Button>
          }
        </Navbar.Collapse>
      </Container>
    </Navbar>
  )
}

export default AppBar