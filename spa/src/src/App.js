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
import Topbar from './portals/TopBar'


class App extends Component {

  constructor(props) {
    super(props);
    this.state = {apimessage: 'Click button to get answer'}
  }

  render() {
    return (
      <BrowserRouter>
      <div className="App">
      <section className="section">
        <Topbar>
          <div style={{background: "aquamarine"}}>
            <Link to="/">
              <span className="app-title">🎨 PAITINGS CORNER</span>
            </Link>
          </div>
        </Topbar>
        
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
