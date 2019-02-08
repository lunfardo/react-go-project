import React, { Component } from 'react';
import { PaintContext } from '../AppContexts'


class PaintProvider extends Component {
    state = {
      paintsList: [],
      submitPaintForm: this.submitPaintForm
    }

    async loadPaintsFromAPI(){
      return fetch(`//${window.location.host}/api/v1/paints`)
      .then((response) => {
        if(response.ok) {
          return response.json()
        } else {
          return 'not ok'
          }
      })
      .catch(function(error) {
        this.setState({apimessage: 'Fetch failed:' + error.message})
      });  
    }

    componentDidMount() {
        this.loadPaintsFromAPI().then((result)=>{
            if(result!=null){
              this.setState({paintsList: result})
            }    
        })          
    }  
    
    async submitPaintForm(event) {
      event.preventDefault()

      //create data object before request
      var data = new FormData()
      data.append('paintFile', this.state.paintFile)
      data.append('Name', this.state.name)      
      data.append('CurrentLocation', this.state.currentLocation)      
      data.append('PainterId', parseInt(this.state.painter))      
      console.log(data)

      const rawResponse = await fetch(`//${window.location.host}/api/v1/paints`, {
        method: 'POST',
        body: data
      });
      
      if(rawResponse.status == 200) {
        this.setState({redirect:true})
      }      
    }

    render() {
      return (
        <PaintContext.Provider value={this.state}>
          {this.props.children}
        </PaintContext.Provider>
      )}
}

export default PaintProvider