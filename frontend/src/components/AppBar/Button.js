import React from 'react'

const Button = ({
  text,
  onClick
}) => {

  return (
    <button className="appbar-button" onClick={onClick}>
      {text}
    </button>
  )
}

export default Button