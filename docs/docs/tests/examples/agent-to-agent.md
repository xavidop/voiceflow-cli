# Agent-to-Agent Testing Examples

This page provides comprehensive examples of Agent-to-Agent tests for various industries and use cases.

## Basic Examples

### Simple Customer Service

```yaml
name: Basic Customer Inquiry
description: Test agent's ability to handle simple customer questions

agent:
  goal: "Get information about store hours and location"
  persona: "A polite customer who needs basic store information"
  maxSteps: 8
  userInformation:
    - name: 'location'
      value: 'downtown'
```

### User Account Issue

```yaml
name: Account Login Problem
description: Test agent's ability to help with login issues

agent:
  goal: "Resolve a login problem and successfully access my account"
  persona: "A frustrated user who can't remember their password"
  maxSteps: 12
  userInformation:
    - name: 'email'
      value: 'user@example.com'
    - name: 'username'
      value: 'johndoe123'
    - name: 'last_login'
      value: '2 weeks ago'
```

## Industry-Specific Examples

### Banking & Financial Services

#### ATM Transaction Support

```yaml
name: ATM Transaction Help
description: Test agent's ability to assist with ATM issues

agent:
  goal: "Complete a cash withdrawal after card was temporarily blocked"
  persona: "An anxious customer whose card stopped working at an ATM"
  maxSteps: 15
  userInformation:
    - name: 'account_number'
      value: '1234567890'
    - name: 'phone'
      value: '555-123-4567'
    - name: 'card_last_four'
      value: '5678'
    - name: 'withdrawal_amount'
      value: '200'
    - name: 'atm_location'
      value: 'Main Street Branch'
```

#### Credit Card Dispute

```yaml
name: Credit Card Dispute Resolution
description: Test agent's ability to handle credit card disputes

agent:
  goal: "Report an unauthorized charge and request a refund"
  persona: "A concerned customer who noticed a fraudulent charge on their statement"
  maxSteps: 20
  userInformation:
    - name: 'card_number'
      value: '**** **** **** 1234'
    - name: 'dispute_amount'
      value: '$89.99'
    - name: 'merchant_name'
      value: 'Unknown Online Store'
    - name: 'transaction_date'
      value: 'January 15, 2025'
```

### E-commerce & Retail

#### Product Return

```yaml
name: Defective Product Return
description: Test agent's ability to process product returns

agent:
  goal: "Return a broken item and get a full refund"
  persona: "A disappointed customer who received a damaged product"
  maxSteps: 14
  userInformation:
    - name: 'order_number'
      value: 'ORD-789456'
    - name: 'email'
      value: 'customer@email.com'
    - name: 'product_name'
      value: 'Bluetooth Wireless Headphones'
    - name: 'issue'
      value: 'right earbud not working'
    - name: 'purchase_date'
      value: '2024-12-20'
```

#### Order Tracking

```yaml
name: Order Status Inquiry
description: Test agent's ability to provide order updates

agent:
  goal: "Find out when my delayed order will arrive"
  persona: "An impatient customer whose order is late for a special occasion"
  maxSteps: 10
  userInformation:
    - name: 'order_number'
      value: 'ORD-555123'
    - name: 'phone'
      value: '555-987-6543'
    - name: 'delivery_address'
      value: '123 Main St, City, State'
    - name: 'expected_date'
      value: 'January 10, 2025'
```

### Healthcare

#### Appointment Scheduling

```yaml
name: Urgent Appointment Request
description: Test agent's ability to schedule medical appointments

agent:
  goal: "Schedule an urgent appointment with a specialist for recurring symptoms"
  persona: "A worried patient experiencing worsening health symptoms"
  maxSteps: 18
  userInformation:
    - name: 'patient_id'
      value: 'PAT-789123'
    - name: 'insurance_provider'
      value: 'Blue Cross Blue Shield'
    - name: 'symptoms'
      value: 'persistent chest pain and shortness of breath'
    - name: 'preferred_doctor'
      value: 'Dr. Johnson (Cardiology)'
    - name: 'availability'
      value: 'weekday afternoons'
```

#### Prescription Refill

```yaml
name: Prescription Refill Request
description: Test agent's ability to handle prescription refills

agent:
  goal: "Refill my blood pressure medication that's running low"
  persona: "An elderly patient who needs help navigating the refill process"
  maxSteps: 12
  userInformation:
    - name: 'prescription_number'
      value: 'RX-456789'
    - name: 'medication_name'
      value: 'Lisinopril 10mg'
    - name: 'pharmacy'
      value: 'CVS Pharmacy on Oak Street'
    - name: 'doctor_name'
      value: 'Dr. Smith'
```

### Travel & Hospitality

#### Hotel Booking

```yaml
name: Hotel Reservation
description: Test agent's ability to handle hotel bookings

agent:
  goal: "Book a hotel room for a business trip next week"
  persona: "A busy business traveler who needs accommodation quickly"
  maxSteps: 16
  userInformation:
    - name: 'loyalty_number'
      value: 'REWARDS123456'
    - name: 'dates'
      value: 'January 20-22, 2025'
    - name: 'destination'
      value: 'Chicago, IL'
    - name: 'room_preference'
      value: 'non-smoking, king bed'
    - name: 'budget'
      value: 'under $200 per night'
```

#### Flight Change Request

```yaml
name: Flight Modification
description: Test agent's ability to handle flight changes

agent:
  goal: "Change my flight to an earlier departure due to a schedule conflict"
  persona: "A stressed traveler whose meeting was moved up unexpectedly"
  maxSteps: 14
  userInformation:
    - name: 'confirmation_number'
      value: 'ABC123'
    - name: 'current_flight'
      value: 'UA456 departing 6:00 PM'
    - name: 'preferred_time'
      value: 'morning departure'
    - name: 'destination'
      value: 'Los Angeles'
    - name: 'travel_date'
      value: 'January 25, 2025'
```

### Technical Support

#### Software Installation Help

```yaml
name: Software Installation Issue
description: Test agent's ability to provide technical support

agent:
  goal: "Get help installing software that keeps failing"
  persona: "A non-technical user who's frustrated with repeated installation failures"
  maxSteps: 20
  userInformation:
    - name: 'operating_system'
      value: 'Windows 11'
    - name: 'software_name'
      value: 'Adobe Creative Suite'
    - name: 'error_message'
      value: 'Installation failed: Error code 1603'
    - name: 'computer_specs'
      value: '16GB RAM, Intel i7 processor'
```

#### Internet Connectivity Issue

```yaml
name: Internet Service Problem
description: Test agent's ability to troubleshoot connectivity issues

agent:
  goal: "Resolve internet outage and restore service to my home"
  persona: "A work-from-home professional whose internet suddenly stopped working"
  maxSteps: 18
  userInformation:
    - name: 'account_number'
      value: 'ISP-789456'
    - name: 'service_address'
      value: '456 Elm Street, Suburb, State'
    - name: 'plan_type'
      value: 'Business 500 Mbps'
    - name: 'equipment'
      value: 'Cisco modem and Netgear router'
    - name: 'outage_duration'
      value: '3 hours'
```

## Complex Scenarios

### Multi-Step Customer Journey

```yaml
name: Complete Customer Onboarding
description: Test agent's ability to guide new customers through full onboarding

agent:
  goal: "Complete account setup, verify identity, and make first transaction"
  persona: "A cautious new customer who wants to understand each step thoroughly"
  maxSteps: 25
  userInformation:
    - name: 'full_name'
      value: 'Sarah Johnson'
    - name: 'email'
      value: 'sarah.johnson@email.com'
    - name: 'phone'
      value: '555-456-7890'
    - name: 'ssn_last_four'
      value: '1234'
    - name: 'address'
      value: '789 Pine Ave, Cityville, State 12345'
    - name: 'employment'
      value: 'Marketing Manager at Tech Corp'
    - name: 'initial_deposit'
      value: '1000'
```

### Crisis Management

```yaml
name: Service Outage Communication
description: Test agent's ability to handle service disruption inquiries

agent:
  goal: "Get updates about the service outage and estimated restoration time"
  persona: "An angry customer whose business is affected by the outage"
  maxSteps: 12
  userInformation:
    - name: 'business_account'
      value: 'BIZ-654321'
    - name: 'service_type'
      value: 'Enterprise Cloud Hosting'
    - name: 'affected_services'
      value: 'email server and website'
    - name: 'business_impact'
      value: 'cannot process customer orders'
```

## Personality Variations

### Different Customer Types

#### The Patient Customer

```yaml
name: Patient Customer Interaction
description: Test with a calm, understanding customer persona

agent:
  goal: "Resolve a billing discrepancy in my monthly statement"
  persona: "A patient, polite customer who understands that mistakes happen"
  maxSteps: 10
  userInformation:
    - name: 'account_number'
      value: 'ACC-112233'
    - name: 'billing_month'
      value: 'December 2024'
    - name: 'disputed_amount'
      value: '$25.99'
```

#### The Urgent Customer

```yaml
name: High-Priority Customer Issue
description: Test with a time-sensitive, urgent customer persona

agent:
  goal: "Get immediate help with a payment that failed during checkout"
  persona: "An urgent customer who needs immediate resolution for a time-sensitive purchase"
  maxSteps: 8
  userInformation:
    - name: 'transaction_id'
      value: 'TXN-998877'
    - name: 'card_ending'
      value: '4567'
    - name: 'purchase_amount'
      value: '$299.99'
    - name: 'deadline'
      value: 'need to complete purchase today'
```

#### The Confused Customer

```yaml
name: Confused Customer Guidance
description: Test with a customer who needs extra explanation

agent:
  goal: "Understand how to use the new mobile app features"
  persona: "A confused customer who is not tech-savvy and needs step-by-step guidance"
  maxSteps: 16
  userInformation:
    - name: 'phone_type'
      value: 'iPhone 12'
    - name: 'app_version'
      value: '2.1.0'
    - name: 'specific_feature'
      value: 'mobile check deposit'
```

## Tips for Effective Agent Tests

### 🎯 Goal Writing Best Practices

- **Be Specific**: Instead of "get help", use "resolve login issue and access account"
- **Include Context**: Mention why the goal is important to the user
- **Make it Measurable**: Define what "success" looks like

### 👤 Persona Development

- **Add Emotion**: Include emotional state (frustrated, confused, urgent)
- **Technical Level**: Specify technical expertise level
- **Background Context**: Provide relevant situational details

### 📊 User Information Strategy

- **Essential Data**: Include information your agent typically needs
- **Realistic Values**: Use believable names, addresses, and IDs
- **Complete Coverage**: Provide all data types your agent might request

### ⚡ Step Optimization

- **Start Conservative**: Begin with fewer steps and increase if needed
- **Monitor Logs**: Check test logs to see actual step usage
- **Buffer for Edge Cases**: Add 20-30% buffer for unexpected paths

These examples demonstrate the flexibility and power of Agent-to-Agent testing across different industries and scenarios. Each test simulates realistic user behavior while working toward specific, measurable goals.
