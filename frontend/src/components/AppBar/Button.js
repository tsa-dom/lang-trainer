import React from 'react'

const Button = ({
  text,
  onClick
}) => {

  return (
    <a className="appbar-button" onClick={onClick}>
      {text}
    </a>
  )
}

export default Button