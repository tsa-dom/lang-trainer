import React from 'react'
import { Container, Button, Navbar, Dropdown, DropdownButton } from 'react-bootstrap'
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
          <DropdownButton style={{ marginRight: 20 }} title={t(language)}>
            {languages.map(r => <Dropdown.Item key={r}>{t(r)}</Dropdown.Item>)}
          </DropdownButton>
          <Button className="button-menu">{t('login')}</Button>
        </Navbar.Collapse>
      </Container>
    </Navbar>
  )
}

export default AppBar