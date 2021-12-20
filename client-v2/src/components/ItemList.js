import React from 'react'
import { Container, Table } from 'react-bootstrap'

// eslint-disable-next-line no-unused-vars
const ItemList = ({ rows, columns, onCellClick, handleItemRemove, colCount = 10 }) => {
  const items = rows.map(row => <tr onClick={() => onCellClick(row)} style={{ cursor: 'pointer' }} key={row.id}>
    {columns.map(col => <td key={col.field}>{row[col.field]}</td>)}
  </tr>)

  const pageCount = Math.ceil(items.length / colCount)

  const pages = rows.map(row => <tr onClick={() => onCellClick(row)} style={{ cursor: 'pointer' }} key={row.id}>
    {columns.map(col => <td key={col.field}>{row[col.field]}</td>)}
  </tr>)

  console.log(pages, pageCount)

  return (
    <Container>
      <Table striped bordered hover >
        <thead>
          <tr>
            {columns.map(col => <th key={col.field}>{col.headerName}</th>)}
          </tr>
        </thead>
        <tbody>
          {items}
        </tbody>
      </Table>
    </Container>
  )
}

export default ItemList