# Context Based Identity-Managment Method

**Abstract**
The idea proposes a new method for mobile and immobile asset protection and secure life cycle management using network assisted and network provided events.
The method to
1. Generate mobility and communication context events.
2. Store communication context event data.
3. Generate a Hash Tag based on stored event data.
4. Authorize group users to acquire hash tags for multiuser authentication.
5. Induced/Virtual roaming for relative events.
6. Transferring the Assets to the next trusted user using mobile network events 

Events can be location events or device state events or mobile context events specific to device or relative to device used to prove the authorized owner 
are specific to these domain 

1.	IoT: Smart City
2.	General Public: Protecting/Monitoring the assets 
3.	Government/Private Body: Protecting/Monitoring the assets
4.	Context Based Resource Authorization and Authentication
5.	Context Based Identity Management.
6.	Supply Chain Controls
7.	Electronic tokens for group asset managments

Some mobility Event Examples
1.	Location Events (Real, Induced)
2.	User state Event (Real, Induced)
3.	Roaming Events (Real, Induced) (includes space roaming)
4.	Re-Attach Event (Real, Induced)
5.	Detach Event (Induced)
6.	Relative Attach Event (Real)
7.	OAM events
8.	Charging Events
9.	Attach from type of network event
10.	Attach from specified Device event
11.	Receive SMS event
12.	Send SMS event 
13.	Subscribe for SMS event.
14.	Generate email event
15.	Get Charging event.
16.	Handover from RAT (inter/intra) events
17.	Handover from different access events(includes satellite access)
18.	Using multiple SIMs from multi-vendor operator related events
19.	Events due to OTT

**Realization**
 ![image](https://github.com/KiranCS-17/identity-managment/blob/main/figure-1.png)
 
  1. E1, E2: Location Events generated by Devices 
  2. S1, S2: Location Events generated by RANs w.r.t source.
  3. S3, S4: Location Event Generation control by cloud core
  4. S5: Event store with specific pattern.

  Vicinity-Vr is the case of roaming by the same device D1 of the CUG.
  
  ##  Sample Message Flow and descriptions
  ![image](https://github.com/KiranCS-17/identity-managment/blob/main/figure-2.png)
  
**Note:**
The method of generating the events, inducing it and maintaining the records of events for UE and group of UEs, is hosted inside an AF called Communication Context & Event  Management Function.( CCEF)
This AF is responsible for producing the Authorization events and creating challenge factors for fraud user using time stamp or any related events as a question.
Some of the concepts of induced virtual roaming with AR-VR facilitation to create mobile communication context events that are encoded in 5G Data stores like UDR/UDSF.
The new application function CCEM (communication context and event Management) is used to trigger mobility events related to asset protection and formulating Event data stored in granular time series DB in compact method with context data chains. 
The Fraud user will be challenged to prove the Mobility Context Events and its relative events, at specified granular time or chain of times that becomes NP hard problem to solve and with multiple hashes the trusted user can request for recursive challenges with the recursino depth configured (medium,high,or greedy) . 
Block chain method will be one implementation method to realize this concept.
Google APIs can be used as an example for Location based Authentication and Authorization apart from 3GPP defined procedures.

1.  UE/CUG performs authentication and attach subscribing for this new monitoring service.
2.  RAN forwards this request to Serving Nodes (MME, AMF).
3.  Serving node asks HSS+UDM to perform authentication authorization.
4.  HSS+UDM reads the subscriber profile in DB.
5.  HSS+UDM gets the subscriber profile response from DB.
    5.1,5.2,5.3 The UE/CUG is authenticated and attached and acknowledged by network.
6.  HSS/UDM trigger CCEF to perform the service execution for this UE/CUG.
7.  CCEF asks HSS to generate the location events, or any other events mentioned in(Type of Events).
8.  HSS/UDM queries the serving node to get the event data.
9.  HSS/UDM gets event data from serving nodes.
10. HSS/UDM forwards the event data to CCEF.
11. CCEF acknowledges the HSS with permutation data.
        At this step the CCEF implicitly subscribes for permutation data movement from UDR to UDSF based on time stamp.
12. HSS/UDM updates the permutation data into the subscriber profile/CUG.
13. The CCEF Executed the Algorithm A1 and stores the output into UDSF in time series DB.
14. The UDSF acknowledges the data.

Based on user query or periodicity the CCEF can produce the Geo Hash with Algorithm A2
And supply to user that can be used as challenge against fraud.

**Goal**: The goal of this project is to realize the CCEF Application Function as PoC in step 1.
          Further steps focuses on performance aspects concerning response times.


**Sample Real Usage Example**
The CCEF instructs the user to perform induced roaming using Virtual reality method.
The user uses Virtual reality headsets and viewers and tries to roam from source (either chosen by network or by user)
and moves to the destination (Geo Asset) using a route map.
This virtually creates a roaming scenario.
In this roaming scenario the specific events can be induced by operator between the resource and CUG and the owner  
by using CCEF to get the location information, user state, the user specific contexts.
The other relative events from control plane elements are recorded as event data stored in UDSF.
The event can be fully constructed based on the induced roaming. 

A hash-code is generated for this with input as Event data and stored against the time stamp.

  

  

