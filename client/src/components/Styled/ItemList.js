import { DataGrid } from '@mui/x-data-grid'
import { Button } from '@material-ui/core'
import React, { useState } from 'react'
import './index.css'

const ItemList = ({
  rows,
  columns,
  onCellClick,
  handleItemRemove
}) => {
  const [selected, setSelected] = useState([])

  if (!rows) return <></>

  return (
    <div style={{ position: 'relative', height: 'calc(100vh - 300px)', width: '100%' }}>
      <div style={{ position: 'absolute', top: 0, width: '100%', height: '100%' }}>
        <DataGrid
          rows={rows}
          columns={columns}
          checkboxSelection
          disableSelectionOnClick
          onCellClick={onCellClick}
          onSelectionModelChange={(value) => setSelected(value)}
          disableExtendRowFullWidth
        />
      </div>
      {handleItemRemove &&
        <div className="item-list-delete-button">
          <Button
            style={{ color: 'red', borderColor: 'red' }}
            variant="outlined"
            onClick={() => handleItemRemove(selected)}
          >
            DELETE
          </Button>
        </div>
      }
    </div>
  )
}

export default ItemList