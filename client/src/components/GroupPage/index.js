import React, { useState } from 'react'
import { useSelector } from 'react-redux'
import MenuBar from '../Styled/MenuBar'

const GroupPage = () => {
  const group = useSelector(state => state.groups.selectedGroup)
  const [selected, setSelected] = useState('word-list')

  return (
    <>
      <div className="page-container-head">
        <div className="page-container-header">{group.name}</div>
        <hr className="page-container-linebreak"></hr>
        <MenuBar
          selected={selected}
          setSelected={setSelected}
          items={['word-list']}
        />
      </div>
      <div className="page-container-body">

      </div>
    </>
  )
}

export default GroupPage