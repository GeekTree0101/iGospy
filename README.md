### [Clean Swift Scaffold is OPEN!](https://github.com/GeekTree0101/clean-swift-scaffold)
<img src="https://github.com/GeekTree0101/clean-swift-scaffold/blob/develop/logo.png" />

iGoSpy will no longer be maintained


# iGoSpy [DEPRECATED]
Clean Swift spy generator built on Go

# Support..
- interactor interface & spy
- presenter interface & spy 
- display interface & spy

# Guide

### setup

```sh
// FIXME
cd app
npm install
npm run build
cd ../
go build
./iGospy run

```

### Input usecase

<img src="https://github.com/GeekTree0101/iGospy/blob/develop/res/guide1.png" />
          
```swift

enum KarrotFeedModels {

  enum Reload {
    
    struct Req {
      
    }
    
    struct Res {
      
      var seq: FeedSequence?
      var error: Error?
    }
    
    struct ViewModel {
      
      var changeSet: [Change<KarrotFeedItemViewModel>]
      var error: Error?
    }
  }
  
  enum Next {
    
    struct Req {
      
    }
    
    struct Res {
      
      var seq: FeedSequence?
      var error: Error?
    }
    
    struct ViewModel {
      
      var changeSet: [Change<KarrotFeedItemViewModel>]
      var error: Error?
    }
  }
}     
```
   
Now you just copy the result. \o/ 

### Interactor output
          
<img src="https://github.com/GeekTree0101/iGospy/blob/develop/res/guide2.png" />
   
### Presenter output   

<img src="https://github.com/GeekTree0101/iGospy/blob/develop/res/guide3.png" />
          
### Display output
         
<img src="https://github.com/GeekTree0101/iGospy/blob/develop/res/guide4.png" />
