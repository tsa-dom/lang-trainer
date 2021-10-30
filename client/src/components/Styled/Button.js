import React from 'react'
import './index.css'

const Button = ({
  text,
  onClick,
  className
}) => {

  return (
    <button type="submit" className={className} onClick={onClick}>
      {text}
    </button>
  )
}

export default Button