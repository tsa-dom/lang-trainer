import { DataGrid } from '@mui/x-data-grid'
import React from 'react'

const ItemList = ({
  rows,
  columns,
  onCellClick
}) => {

  return (
    <div style={{ height: 'calc(100vh - 300px)', width: '100%' }}>
      <DataGrid
        rows={rows}
        columns={columns}
        checkboxSelection
        disableSelectionOnClick
        onCellClick={onCellClick}
      />
    </div>
  )
}

export default ItemList