import React from 'react'

const Input = ({
  label,
  className,
  type,
  onChange,
  id,
  value
}) => {

  return (
    <>
      {label &&
        <div className={`${className}-label`}>{label}</div>
      }
      <input
        id={id}
        className={`${className}-field`}
        type={type}
        onChange={onChange}
        value={value}
      />
    </>
  )
}

export default Input