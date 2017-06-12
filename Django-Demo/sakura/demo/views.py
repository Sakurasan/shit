from django.shortcuts import render
from django.http import HttpResponse
from . import models
from django.http import HttpResponseRedirect  

# Create your views here.


def index(request):
    # return HttpResponse('Hello world')
    # return render(request,'index.html',{'hello':'Hello My Django Web'})
    
    articles = models.Artical.objects.all()
    return render(request,'index.html',{'articles':articles})

def article_page(request,article_id):
    article = models.Artical.objects.get(pk=article_id)
    return render(request,'demo/article.html',{'article':article})

def edit_page(request,article_id):
    if str(article_id) == '0':
        return render(request,'demo/edit.html')
    article = models.Artical.objects.get(pk=article_id)

    return render(request,'demo/edit.html',{'article':article})

def edit_action(request):
    title = request.POST.get('title','TITLE')
    content = request.POST.get('content','CONTENT')
    
    article_id = request.POST.get('id','0')
    if article_id=='0':
        models.Artical.objects.create(title=title,content=content)
        return HttpResponseRedirect('/blog/')

    article = models.Artical.objects.get(pk=article_id)
    article.title=title
    article.content=content
    article.save()
    # return render(request,'index.html',{'articles':articles})
    return HttpResponseRedirect('/blog/')