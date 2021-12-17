import React from 'react'
import { Container, Button, Navbar, NavDropdown } from 'react-bootstrap'
import { useTranslation } from 'react-i18next'
import { useSelector } from 'react-redux'
import { resources } from '../../i18n'

const AppBar = () => {
  const { t } = useTranslation()
  const language = useSelector(state => state.users.language)

  const languages = Object.keys(resources).map(key => t(resources[key].language))

  return (
    <Navbar className="nav-menu" bg="dark" variant="dark" expand="lg">
      <Container>
        <Navbar.Brand><span className="nav-menu-brand">Lang Trainer</span></Navbar.Brand>
        <Navbar.Toggle />
        <Navbar.Collapse className="justify-content-end">
          <NavDropdown title={t(language)}>
            {languages.map(r => <NavDropdown.Item key={r}>{t(r)}</NavDropdown.Item>)}
          </NavDropdown>
          <NavDropdown style={{ marginRight: 20 }} title={t('groups')}>
            <NavDropdown.Item>{t('group-list')}</NavDropdown.Item>
            <NavDropdown.Item>{t('templates')}</NavDropdown.Item>
          </NavDropdown>
          <Button className="button-menu">{t('login')}</Button>
        </Navbar.Collapse>
      </Container>
    </Navbar>
  )
}

export default AppBar