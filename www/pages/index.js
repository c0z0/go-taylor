import React, { useState } from 'react'
import Head from 'next/head'

export default function Index() {
  const [inputState, setInputState] = useState('')
  const [functionState, setFunctionState] = useState('exp')
  const [resultState, setResultState] = useState('No result')

  async function onInputChange(input, func) {
    if (input !== inputState) setInputState(input)

    if (func !== functionState) setFunctionState(func)

    const res = await fetch(`/api/calculate.go?command=${func}&input=${input}`)

    const data = await res.json()

    setResultState(data.result)
  }

  return (
    <div className="container">
      <Head>
        <title>Taylor series calculator</title>
      </Head>
      <h1>Taylor series calculator</h1>
      <select
        value={functionState}
        onChange={({ target: { value } }) => onInputChange(inputState, value)}
      >
        <option value="exp">e^x</option>
        <option value="sin">sin x</option>
        <option value="cos">cos x</option>
        <option value="tan">tan x</option>
        <option value="cot">cot x</option>
      </select>
      <input
        value={inputState}
        onChange={({ target: { value } }) =>
          onInputChange(value, functionState)
        }
        placeholder="Input"
      />
      <p>{inputState.length > 0 ? resultState : 'No result'}</p>
      <a href="https://github.com/c0z0/go-taylor">[src]</a>
      <a href="https://cserdean.com">cserdean.com</a>

      <style jsx>{`
        h1 {
          font-weight: 500;
          margin-bottom: 32px;
        }

        .container {
          height: 100vh;
          display: flex;
          flex-direction: column;
          align-items: center;
          justify-content: center;
        }

        input,
        select {
          height: 30px;

          border: none;
          border-bottom: #ddd solid 1px;
          outline: none;
          min-width: 50%;
          width: 300px;
          padding: 12px 20px;
          font-size: 16px;
        }

        select {
          min-width: calc(50% + 40px);
          width: 340px;
        }

        p {
          font-size: 24px;
          min-width: 40%;
          padding: 12px 20px;
          width: 200px;
          text-align: center;
          background: #f2f2f2;
          border-radius: 12px;
        }

        a {
          color: #aaa;
          transition: all 0.2s;
          margin-bottom: 1em;
        }

        a:hover {
          color: #222;
        }
      `}</style>
      <style jsx global>{`
        body {
          font-size: 16px;
          min-height: 100%;
          position: relative;
          font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto,
            Helvetica, Arial, sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji',
            'Segoe UI Symbol';
          margin: 0;
          color: #222;
          padding: 0;
        }

        html {
          padding: 0;
        }
      `}</style>
    </div>
  )
}
