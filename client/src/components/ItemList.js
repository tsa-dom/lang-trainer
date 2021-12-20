/* eslint-disable no-unused-vars */
import React, { useState } from 'react'
import { Button, Container, Table } from 'react-bootstrap'
import { Checkbox, Text } from '@chakra-ui/react'
import { useTranslation } from 'react-i18next'

const ItemList = ({ rows, columns, onCellClick, handleItemRemove, title }) => {
  const { t } = useTranslation()
  const [checked, setChecked] = useState([])

  const handleItemCheck = (e, id) => {
    e.target.checked ? setChecked(checked.concat(id)) : setChecked(checked.filter(c => c !== id))
  }

  const getIds = () => rows.map(r => r.id)

  const items = rows.map(row => <tr key={row.id}>
    <td style={{ width: 0 }}>
      <Checkbox
        isChecked={checked.includes(row.id)}
        onChange={e => handleItemCheck(e, row.id)}
        colorScheme='blue'
        style={{ marginLeft: 2, marginRight: 2, marginTop: 4, borderColor: 'grey' }}
      />
    </td>
    {columns.map(col => <td
      key={col.field}
      style={{ cursor: 'pointer', width: col.width }}
      onClick={() => onCellClick(row)}
    >{row[col.field]}</td>)}
  </tr>)

  return (
    <Container style={{ position: 'relative' }}>
      <Text fontSize='4xl' style={{ marginTop: 10 }}>{title}</Text>
      <Table striped bordered hover >
        <thead>
          <tr>
            <th>
              <Checkbox
                isChecked={checked.length === rows.length}
                isIndeterminate={checked.length > 0 && checked.length < rows.length}
                colorScheme='blue'
                style={{ marginLeft: 2, marginRight: 2, marginTop: 4, borderColor: 'grey' }}
                onChange={() => checked.length ? setChecked([]) : setChecked(getIds())}
              />
            </th>
            {columns.map(col => <th key={col.field}>{col.headerName}</th>)}
          </tr>
        </thead>
        <tbody>
          {items}
        </tbody>
      </Table>
      <Button
        style={{ position: 'absolute', right: 0, backgroundColor: 'red', borderColor: 'red' }}
        onClick={() => {
          handleItemRemove(checked)
          setChecked([])
        }}
      >
        {t('delete')}
      </Button>
    </Container>
  )
}

export default ItemList