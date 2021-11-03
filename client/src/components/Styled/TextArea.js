import React from 'react'

const TextArea = ({
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
      <textarea
        id={id}
        className={`${className}-field`}
        type={type}
        onChange={onChange}
        value={value}
      />
    </>
  )
}

export default TextArea