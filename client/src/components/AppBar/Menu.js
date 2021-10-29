import React from 'react'
import Button from './Button'

const Menu = ({ setPage, showMenu }) => {

  const first = () => {
    setPage('first')
  }

  const second = () => {
    setPage('second')
  }

  const third = () => {
    setPage('third')
  }

  const fourth = () => {
    setPage('fourth')
  }

  if (!showMenu) return <div className="appbar-menu"></div>



  return (
    <div className="appbar-menu">
      <Button text="Button 1" onClick={first} />
      <Button text="Button 2" onClick={second} />
      <Button text="Button 3" onClick={third} />
      <Button text="Button 4" onClick={fourth} />
    </div>
  )
}

export default Menu