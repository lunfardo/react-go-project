import React, { Component } from 'react';
import './bulma.min.css';
import './custom.css';
import { BrowserRouter, Route, Link } from 'react-router-dom'
import PainterProvider from './providers/PainterProvider'
import PaintProvider from './providers/PaintProvider'
import PaintersList from './components/painters/PaintersList'
import PaintsList from './components/paints/PaintsList'
import PaintForm from './components/paints/PaintForm'
import PainterShow from './components/painters/PainterShow'
import PainterForm from './components/painters/PainterForm'


class App extends Component {

  constructor(props) {
    super(props);
    this.state = {apimessage: 'Click button to get answer'}
  }

  render() {
    return (
      <BrowserRouter>
      <div className="App">
      <section class="section">
        <Link to="/">
          <span class="app-title">The Paints Corner</span>
        </Link>
        
          <PainterProvider>
          <PaintProvider>

            <Route path="/" exact component={PaintersList}/>
            <Route path="/painter/:painterId" exact component={PainterShow}/>
            <Route path="/painter/:painterId/paints" exact component={PaintsList}/>
            <Route path="/painter/create" exact component={PainterForm}/>
            <Route path="/paint/create" exact component={PaintForm}/>

          </PaintProvider>
          </PainterProvider>
        </section>
    </div>
    </BrowserRouter>      
    );
  }
}

export default App;
