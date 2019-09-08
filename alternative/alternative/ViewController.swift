//
//  ViewController.swift
//  MRTowerDiffence
//
//  Created by 洞井僚太 on 2019/08/24.
//  Copyright © 2019 洞井僚太. All rights reserved.
//

import UIKit
import SceneKit
import ARKit
import SpriteKit
import GameplayKit
//convenience init(string: Any?, extrusionDepth: CGFloat)
class ViewController: UIViewController, ARSCNViewDelegate,SCNPhysicsContactDelegate {
    var enemies:[SCNNode]=[]
    @IBOutlet var sceneView: ARSCNView!
    var score = 0
   

    
    var bullets:[SCNNode] = []
    let bulletCategory = 0x0000
    let enemyCategory = 0x0001
    
    var spawnTimer:Timer!
    var updateTimer:Timer!
    var counter = 0
    var theta = 0.0
    var str:SCNText!
    var textNode:SCNNode!
    //let scoreLabelView:
    override func viewDidLoad() {
        super.viewDidLoad()
        str = SCNText(string: "score:\(score)", extrusionDepth: 0)
        str.font = UIFont(name: "PixelMplus12-Bold", size: 5)
        textNode = SCNNode(geometry: str)
       // scoreLabel = UILabel()
  //      scoreLabel.text = "score:\(score)"
    //    scoreLabel.font = UIFont.systemFont(ofSize: 30.0)
     /*   scoreLabel.frame.origin.x = view.frame.width/2-20
        scoreLabel.frame.origin.y = view.frame.height/2
 view.addSubview(scoreLabel)*/
        
        //scoreLabel.font
       /* let scene = GameScene()
        let scnView = self.view as! SCNView
        scnView.backgroundColor = UIColor.black
        scnView.scene = scene
        scnView.showsStatistics = true
 scnView.allowsCameraControl = true*/
        // Set the view's delegate
        // Show statistics such as fps and timing information
        sceneView.showsStatistics = true
        spawnTimer = Timer.scheduledTimer(withTimeInterval: 2, repeats: true, block: {_ in self.addEnemy()})
        spawnTimer = Timer.scheduledTimer(withTimeInterval: 0.1, repeats: true, block: {_ in self.gameover()})
        // Create a new scene
        let scene = SCNScene(named: "art.scnassets/enemy_red_walk2.scn")!
        
        // Set the scene to the view
        sceneView.scene = scene
        sceneView.delegate = self
        sceneView.scene.physicsWorld.contactDelegate = self
         textNode.position = SCNVector3(0, 10, -30)
        sceneView.scene.rootNode.addChildNode(textNode)
        
    }
    func addEnemy(){
        print("kokoniiru")
        let box:SCNGeometry = SCNBox(width: 4, height: 4, length: 4, chamferRadius: 0)
        let physicsshape = SCNPhysicsShape(geometry: box, options: nil)
        var enemy = SCNNode(geometry: box)
       // enemy.geometry = box
        enemy.physicsBody = SCNPhysicsBody(type: .dynamic, shape: physicsshape)
        let material = SCNMaterial()
        material.diffuse.contents = UIColor.white
        enemy.geometry?.materials = [material]
        enemy.physicsBody?.categoryBitMask = enemyCategory
        enemy.physicsBody?.collisionBitMask = bulletCategory
        enemy.physicsBody?.contactTestBitMask = bulletCategory
        enemy.position = SCNVector3Make(100*Float(sin(theta/Double.pi)),0,-100*Float(cos(theta/Double.pi)))
        enemies.append(enemy)
        let targetPos = SCNVector3(0,0,0)
       //let target = camera.convertPosition(targetPos, to: nil)
        let action = SCNAction.move(to: targetPos, duration: 20)
        //let remove = SCNAction.removeFromParentNode()
        enemy.runAction(action)
        sceneView.scene.rootNode.addChildNode(enemy)
       /* let sceneSource = SCNSceneSource(url: url, options: nil)!
        guard let modelNode = sceneSource.entryWithIdentifier("enemy_red_walk2", withClass: SCNNode.self) else {
            fatalError()
        }*/
        counter += 10
        theta += 10
    }
    func gameover(){
        textNode.removeFromParentNode()
        str = SCNText(string: "score:\(score)", extrusionDepth: 10)
        str.font = UIFont(name: "PixelMplus12-Bold", size: 5)
        textNode = SCNNode(geometry: str)
        textNode.position = SCNVector3(0, 10, -30)
        self.sceneView.scene.rootNode.addChildNode(textNode)
        if enemies.count>0{
        for i in 0..<enemies.count-1{
            if enemies[i].position.x == 0 && enemies[i].position.z == 0 {
                let str = SCNText(string: "Game Over", extrusionDepth: 0)
                str.font = UIFont(name: "PixelMplus12-Bold", size: 5)
                let gameoverNode = SCNNode(geometry: str)
                gameoverNode.position = SCNVector3Make(0, 0, -30)
                sceneView.scene.rootNode.addChildNode(gameoverNode)
                sceneView.scene.isPaused = true
            }
        }
        }
    }
    override func viewWillAppear(_ animated: Bool) {
        super.viewWillAppear(animated)
        
        // Create a session configuration
        let configuration = ARWorldTrackingConfiguration()
        
        // Run the view's session
        sceneView.session.run(configuration)
    }
    
    override func viewWillDisappear(_ animated: Bool) {
        super.viewWillDisappear(animated)
        
        // Pause the view's session
        sceneView.session.pause()
    }
    override func touchesBegan(_ touches: Set<UITouch>, with event: UIEvent?) {
      //  let scnView = self.view as! SCNView
        
     //   scnView.scene?.rootNode.addChildNode(ballNode.clone())
        guard let camera = sceneView.pointOfView else {
            return
        }
        let sphere:SCNGeometry = SCNSphere(radius: 0.5)
        let physicsshape = SCNPhysicsShape(geometry: sphere, options: nil)
        var bullet = SCNNode(geometry: sphere)
        // enemy.geometry = box
        bullet.physicsBody = SCNPhysicsBody(type: .dynamic, shape: physicsshape)
       // var bullet = SCNNode()
       // bullet.geometry = SCNSphere(radius: 0.5)
        let material = SCNMaterial()
        material.diffuse.contents = UIColor.brown
        bullet.geometry?.materials = [material]
        bullet.position = camera.position
        bullet.physicsBody?.categoryBitMask = bulletCategory
        bullet.physicsBody?.collisionBitMask = enemyCategory
        bullet.physicsBody?.contactTestBitMask = enemyCategory
        bullets.append(bullet)
        //let scene = SCNScene(named: "art.scnassets/ship.scn")!
        
        let targetPos = SCNVector3Make(0,0,-100)
        let target = camera.convertPosition(targetPos, to: nil)
        let action = SCNAction.move(to: target, duration: 1)
        let remove = SCNAction.removeFromParentNode()
        bullet.runAction(SCNAction.sequence([action,remove]))
        sceneView.scene.rootNode.addChildNode(bullet)
    }
    func physicsWorld(_ world: SCNPhysicsWorld, didBegin contact: SCNPhysicsContact) {
        score += 1000
        let nodeA = contact.nodeA
        let nodeB = contact.nodeB
        var enemy = SCNNode()
        var bullet = SCNNode()
        if nodeA.categoryBitMask == bulletCategory{
            enemy = nodeB
            bullet = nodeA
        }else{
            enemy = nodeA
            bullet = nodeB
        }
     //   print(score)
        for i in 0..<bullets.count{
            if bullets.count <= i{
                break
            }
            if bullet.position.z == bullets[i].position.z{
                bullets[i].removeFromParentNode()
                bullets.remove(at:i)
            }
        }
        print(enemies.count)
        for i in 0..<enemies.count-1{
          //  print(i,enemies[i])
            if enemy.position.x == enemies[i].position.x{
                enemies[i].removeFromParentNode()
                enemies.remove(at:i)
            }
            
        }
        bullet.removeFromParentNode()
    }
    // MARK: - ARSCNViewDelegate
    
    /*
     // Override to create and configure nodes for anchors added to the view's session.
     func renderer(_ renderer: SCNSceneRenderer, nodeFor anchor: ARAnchor) -> SCNNode? {
     let node = SCNNode()
     
     return node
     }
     */  /*func gameover(){
     for enemy in enemies{
     //if (enemy.position == SCNVector3(x:0,y:0,z:0))
     {
     
     }
     }
     }*/
    /* func hitTest(point:CGPoint,types:ARHitTestResult.ResultType)->[ARHitTestResult]{
     
     }*/
    func session(_ session: ARSession, didFailWithError error: Error) {
        // Present an error message to the user
        
    }
    
    func sessionWasInterrupted(_ session: ARSession) {
        // Inform the user that the session has been interrupted, for example, by presenting an overlay
        
    }
    
    func sessionInterruptionEnded(_ session: ARSession) {
        // Reset tracking and/or remove existing anchors if consistent tracking is required
        
    }
    
}
