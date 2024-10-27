import React, { useState } from 'react'
import { Card, CardContent, CardHeader, CardTitle, CardDescription } from "@/components/ui/card"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { Badge } from "@/components/ui/badge"
import { X, Building2, Users, MapPin, AlertCircle, Plus, CheckCircle } from "lucide-react"
import { Alert, AlertDescription } from "@/components/ui/alert"
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs"

export default function DistributorDashboard() {
  const [formData, setFormData] = useState({
    name: '',
    parent: '',
    include: [],
    exclude: []
  })
  const [includeInput, setIncludeInput] = useState('')
  const [excludeInput, setExcludeInput] = useState('')
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState('')
  const [apiResponse, setApiResponse] = useState(null)

  const [checkForm, setCheckForm] = useState({
    name: '',
    region: ''
  })
  const [checkLoading, setCheckLoading] = useState(false)
  const [checkError, setCheckError] = useState('')
  const [checkResponse, setCheckResponse] = useState(null)

  const handleSubmit = async (e) => {
    e.preventDefault()
    if (!formData.name.trim()) {
      setError('Distributor name is required.')
      return
    }
    setLoading(true)
    setError('')
    setApiResponse(null)
    try {
      const response = await fetch('http://localhost:8000/create-distributor', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(formData),
      })
      const data = await response.json()
      console.log('API Response:', data)
      if (!response.ok) throw new Error(data.message || 'Failed to create distributor')
      setApiResponse(data)
      setFormData({ name: '', parent: '', include: [], exclude: [] })
      setIncludeInput('')
      setExcludeInput('')
    } catch (error) {
      console.error('Error:', error)
      setError(error.message)
    } finally {
      setLoading(false)
    }
  }

  const handleAddRegion = (type) => {
    const input = type === 'include' ? includeInput : excludeInput
    if (!input.trim()) return
    setFormData(prev => ({
      ...prev,
      [type]: [...prev[type], input.trim()]
    }))
    type === 'include' ? setIncludeInput('') : setExcludeInput('')
  }

  const handleRemoveRegion = (region, type) => {
    setFormData(prev => ({
      ...prev,
      [type]: prev[type].filter(r => r !== region)
    }))
  }

  const handleCheckPermission = async (e) => {
    e.preventDefault()
    if (!checkForm.name.trim() || !checkForm.region.trim()) {
      setCheckError('Both distributor name and region are required.')
      return
    }
    
    setCheckLoading(true)
    setCheckError('')
    setCheckResponse(null)
    
    try {
      const response = await fetch('http://localhost:8000/check-permissions', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(checkForm),
      })
      const data = await response.json()
      if (!response.ok) throw new Error(data.message || 'Failed to check permission')
      setCheckResponse(data)
    } catch (error) {
      console.error('Error:', error)
      setCheckError(error.message)
    } finally {
      setCheckLoading(false)
    }
  }

  const handleKeyPress = (e, type) => {
    if (e.key === 'Enter') {
      e.preventDefault()
      handleAddRegion(type)
    }
  }

  return (
    <div className="container mx-auto p-6 max-w-4xl">
      <Tabs defaultValue="create" className="space-y-6">
        <TabsList className="grid w-full grid-cols-2">
          <TabsTrigger value="create">Create Distributor</TabsTrigger>
          <TabsTrigger value="check">Check Permission</TabsTrigger>
        </TabsList>

        <TabsContent value="create">
          <Card className="shadow-lg border-slate-200">
            <CardHeader className="space-y-3 pb-6">
              <CardTitle className="text-xl font-bold text-center text-slate-900">
                Create New Distributor
              </CardTitle>
              <CardDescription className="text-center text-slate-600">
                Fill in the details below to create a new distributor and define their regions
              </CardDescription>
            </CardHeader>
            <CardContent>
              {error && (
                <Alert variant="destructive" className="mb-6">
                  <AlertCircle className="h-4 w-4" />
                  <AlertDescription>{error}</AlertDescription>
                </Alert>
              )}
              {apiResponse && (
              <Alert 
                variant={apiResponse.status === "SUCCESS" ? "border-green-200 text-green-800" : "destructive"} 
                className="mb-6"
              >
                {apiResponse.status === "SUCCESS" ? (
                  <CheckCircle className="h-4 w-4" />
                ) : (
                  <AlertCircle className="h-4 w-4" />
                )}
                <AlertDescription>{apiResponse.message}</AlertDescription>
              </Alert>
            )}
              <form onSubmit={handleSubmit} className="space-y-6">
                <div className="space-y-4">
                  <div className="relative">
                    <Label htmlFor="name" className="text-sm font-medium text-slate-700 flex items-center gap-2">
                      <Building2 className="h-4 w-4" />
                      Distributor Name
                    </Label>
                    <Input
                      id="name"
                      value={formData.name}
                      onChange={(e) => setFormData(prev => ({ ...prev, name: e.target.value }))}
                      placeholder="Enter distributor name"
                      className="mt-1.5"
                      required
                    />
                  </div>

                  <div className="relative">
                    <Label htmlFor="parent" className="text-sm font-medium text-slate-700 flex items-center gap-2">
                      <Users className="h-4 w-4" />
                      Parent Distributor
                    </Label>
                    <Input
                      id="parent"
                      value={formData.parent}
                      onChange={(e) => setFormData(prev => ({ ...prev, parent: e.target.value }))}
                      placeholder="Enter parent distributor"
                      className="mt-1.5"
                    />
                  </div>
                </div>

                <div className="space-y-4">
                  <div className="space-y-2">
                    <Label className="text-sm font-medium text-slate-700 flex items-center gap-2">
                      <MapPin className="h-4 w-4" />
                      Include Regions
                    </Label>
                    <div className="flex gap-2">
                      <Input
                        value={includeInput}
                        onChange={(e) => setIncludeInput(e.target.value)}
                        onKeyPress={(e) => handleKeyPress(e, 'include')}
                        placeholder="Enter region to include (e.g., CENAI-TN-IN)"
                        className="flex-grow"
                      />
                      <Button 
                        type="button" 
                        onClick={() => handleAddRegion('include')}
                        className="flex-shrink-0"
                      >
                        <Plus className="h-4 w-4 mr-1" />
                        Add
                      </Button>
                    </div>
                    <div className="flex flex-wrap gap-1.5 min-h-[40px] p-2 bg-slate-50 rounded-lg">
                      {formData.include.map((region) => (
                        <Badge 
                          key={region} 
                          variant="secondary" 
                          className="px-2 py-1 bg-blue-50 text-blue-700 hover:bg-blue-100 transition-colors"
                        >
                          {region}
                          <X
                            className="ml-1.5 h-3.5 w-3.5 cursor-pointer hover:text-red-500 transition-colors"
                            onClick={() => handleRemoveRegion(region, 'include')}
                          />
                        </Badge>
                      ))}
                    </div>
                  </div>

                  <div className="space-y-2">
                    <Label className="text-sm font-medium text-slate-700 flex items-center gap-2">
                      <MapPin className="h-4 w-4" />
                      Exclude Regions
                    </Label>
                    <div className="flex gap-2">
                      <Input
                        value={excludeInput}
                        onChange={(e) => setExcludeInput(e.target.value)}
                        onKeyPress={(e) => handleKeyPress(e, 'exclude')}
                        placeholder="Enter region to exclude (e.g., CENAI-TN-IN)"
                        className="flex-grow"
                      />
                      <Button 
                        type="button" 
                        onClick={() => handleAddRegion('exclude')}
                        className="flex-shrink-0"
                      >
                        <Plus className="h-4 w-4 mr-1" />
                        Add
                      </Button>
                    </div>
                    <div className="flex flex-wrap gap-1.5 min-h-[40px] p-2 bg-slate-50 rounded-lg">
                      {formData.exclude.map((region) => (
                        <Badge 
                          key={region} 
                          variant="secondary" 
                          className="px-2 py-1 bg-red-50 text-red-700 hover:bg-red-100 transition-colors"
                        >
                          {region}
                          <X
                            className="ml-1.5 h-3.5 w-3.5 cursor-pointer hover:text-red-500 transition-colors"
                            onClick={() => handleRemoveRegion(region, 'exclude')}
                          />
                        </Badge>
                      ))}
                    </div>
                  </div>
                </div>

                <Button 
                  type="submit" 
                  className="w-full text-base py-5"
                  disabled={loading}
                >
                  {loading ? 'Creating Distributor...' : 'Create Distributor'}
                </Button>
              </form>
            </CardContent>
          </Card>
        </TabsContent>

        <TabsContent value="check">
          <Card className="shadow-lg border-slate-200">
            <CardHeader className="space-y-3 pb-6">
              <CardTitle className="text-xl font-bold text-center text-slate-900">
                Check Distribution Permission
              </CardTitle>
              <CardDescription className="text-center text-slate-600">
                Verify if a distributor has permission for a specific region
              </CardDescription>
            </CardHeader>
            <CardContent>
              {checkError && (
                <Alert variant="destructive" className="mb-6">
                  <AlertCircle className="h-4 w-4" />
                  <AlertDescription>{checkError}</AlertDescription>
                </Alert>
              )}
              {checkResponse && (
                <Alert 
                  variant={checkResponse.message === "YES" ? "default" : "destructive"} 
                  className="mb-6"
                >
                  <AlertCircle className="h-4 w-4" />
                  <AlertDescription>
                    Permission {checkResponse.message === "YES" ? "Granted" : "Denied"} for {checkForm.region}
                  </AlertDescription>
                </Alert>
              )}
              <form onSubmit={handleCheckPermission} className="space-y-6">
                <div className="space-y-4">
                  <div className="relative">
                    <Label htmlFor="check-name" className="text-sm font-medium text-slate-700 flex items-center gap-2">
                      <Building2 className="h-4 w-4" />
                      Distributor Name
                    </Label>
                    <Input
                      id="check-name"
                      value={checkForm.name}
                      onChange={(e) => setCheckForm(prev => ({ ...prev, name: e.target.value }))}
                      placeholder="Enter distributor name"
                      className="mt-1.5"
                      required
                    />
                  </div>

                  <div className="relative">
                    <Label htmlFor="check-region" className="text-sm font-medium text-slate-700 flex items-center gap-2">
                      <MapPin className="h-4 w-4" />
                      Region to Check
                    </Label>
                    <Input
                      id="check-region"
                      value={checkForm.region}
                      onChange={(e) => setCheckForm(prev => ({ ...prev, region: e.target.value }))}
                      placeholder="Enter region (e.g., CENAI-TN-IN)"
                      className="mt-1.5"
                      required
                    />
                  </div>
                </div>

                <Button 
                  type="submit" 
                  className="w-full text-base py-5"
                  disabled={checkLoading}
                >
                  {checkLoading ? 'Checking Permission...' : 'Check Permission'}
                </Button>
              </form>
            </CardContent>
          </Card>
        </TabsContent>
      </Tabs>
    </div>
  )
}