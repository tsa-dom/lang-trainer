import React from 'react'
import { Container, Button, Navbar, Dropdown, DropdownButton } from 'react-bootstrap'
import { useTranslation } from 'react-i18next'
import { useSelector } from 'react-redux'
//import { resources } from '../../i18next'

const AppBar = () => {
  const { t } = useTranslation()
  const language = useSelector(state => state.users.language)
  console.log(language)

  //const languages = Object.keys(resources).map(key => t(resources[key].language))
  //console.log(languages)

  return (
    <Navbar className="nav-menu" bg="dark" variant="dark" expand="lg">
      <Container>
        <Navbar.Brand><span className="nav-menu-brand">Lang Trainer</span></Navbar.Brand>
        <Navbar.Toggle />
        <Navbar.Collapse className="justify-content-end">
          <DropdownButton id="dropdown-basic-button" title="Dropdown button">
            <Dropdown.Item href="#/action-1">Action</Dropdown.Item>
            <Dropdown.Item href="#/action-2">Another action</Dropdown.Item>
            <Dropdown.Item href="#/action-3">Something else</Dropdown.Item>
          </DropdownButton>
          <Button className="button-menu">{t('login')}</Button>
        </Navbar.Collapse>
      </Container>
    </Navbar>
  )
}

export default AppBar