import React from 'react'
import './index.css'

const Button = ({
  text,
  onClick,
  className,
  style
}) => {

  return (
    <button
      type="submit"
      className={className ? className : 'default-button'}
      onClick={onClick}
      style={style}
    >
      {text}
    </button>
  )
}

export default Button